// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateMmsDataSourceResponseBodyData) *CreateMmsDataSourceResponseBody
	GetData() *CreateMmsDataSourceResponseBodyData
	SetRequestId(v string) *CreateMmsDataSourceResponseBody
	GetRequestId() *string
}

type CreateMmsDataSourceResponseBody struct {
	Data *CreateMmsDataSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// B42CA730-8187-50F1-9FE0-6733297036DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMmsDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMmsDataSourceResponseBody) GetData() *CreateMmsDataSourceResponseBodyData {
	return s.Data
}

func (s *CreateMmsDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMmsDataSourceResponseBody) SetData(v *CreateMmsDataSourceResponseBodyData) *CreateMmsDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateMmsDataSourceResponseBody) SetRequestId(v string) *CreateMmsDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMmsDataSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMmsDataSourceResponseBodyData struct {
	// example:
	//
	// 18
	DataSourceId *int64 `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
}

func (s CreateMmsDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMmsDataSourceResponseBodyData) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *CreateMmsDataSourceResponseBodyData) SetDataSourceId(v int64) *CreateMmsDataSourceResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *CreateMmsDataSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
