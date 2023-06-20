package kanban

import (
	"path"
	"project-manager-go/data"
)

type KanbanDataProvider struct {
	folder string
}

func NewKanbanDataProvider(folder string) *KanbanDataProvider {
	return &KanbanDataProvider{folder}
}

func (p KanbanDataProvider) Up(ctx *data.DBContext) {
	data.InitializeDemodata[KanbanColumn](ctx, path.Join(p.folder, "kanban_columns.json"))
}

func (p KanbanDataProvider) Down(ctx *data.DBContext) {
	err := ctx.DB.Delete(&KanbanColumn{}, "1 = 1").Error
	if err != nil {
		panic(err)
	}
}
