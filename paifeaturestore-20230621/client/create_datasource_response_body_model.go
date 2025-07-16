// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceId(v string) *CreateDatasourceResponseBody
	GetDatasourceId() *string
	SetRequestId(v string) *CreateDatasourceResponseBody
	GetRequestId() *string
}

type CreateDatasourceResponseBody struct {
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// 1C5B1511-8A5B-59C3-90AF-513F9210E882
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasourceResponseBody) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *CreateDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasourceResponseBody) SetDatasourceId(v string) *CreateDatasourceResponseBody {
	s.DatasourceId = &v
	return s
}

func (s *CreateDatasourceResponseBody) SetRequestId(v string) *CreateDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}
