// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCmsExporterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteCmsExporterResponseBody
	GetData() *string
	SetRequestId(v string) *DeleteCmsExporterResponseBody
	GetRequestId() *string
}

type DeleteCmsExporterResponseBody struct {
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 27E653FA-5958-45BE-8AA9-14D884DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCmsExporterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCmsExporterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCmsExporterResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCmsExporterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCmsExporterResponseBody) SetData(v string) *DeleteCmsExporterResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCmsExporterResponseBody) SetRequestId(v string) *DeleteCmsExporterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCmsExporterResponseBody) Validate() error {
	return dara.Validate(s)
}
