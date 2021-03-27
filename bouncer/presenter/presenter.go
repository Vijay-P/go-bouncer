package presenter

import (
	"io"

	"github.com/wagoodman/go-bouncer/bouncer"
	"github.com/wagoodman/go-bouncer/bouncer/presenter/csv"
	"github.com/wagoodman/go-bouncer/bouncer/presenter/json"
	"github.com/wagoodman/go-bouncer/bouncer/presenter/text"
)

type Presenter interface {
	Present(io.Writer) error
}

func GetPresenter(option Option, results []bouncer.LicenseResult) Presenter {
	switch option {
	case CSVPresenter:
		return csv.NewPresenter(results)
	case JSONPresenter:
		return json.NewPresenter(results)
	case TextPresenter:
		return text.NewPresenter(results)

	default:
		return nil
	}
}
