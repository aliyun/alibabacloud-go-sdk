// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateDatasourceResponseBody
	GetCode() *int32
	SetData(v *CreateDatasourceResponseBodyData) *CreateDatasourceResponseBody
	GetData() *CreateDatasourceResponseBodyData
	SetMessage(v string) *CreateDatasourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDatasourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDatasourceResponseBody
	GetSuccess() *bool
}

type CreateDatasourceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *CreateDatasourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateDatasourceResponseBody) GetData() *CreateDatasourceResponseBodyData {
	return s.Data
}

func (s *CreateDatasourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDatasourceResponseBody) SetCode(v int32) *CreateDatasourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDatasourceResponseBody) SetData(v *CreateDatasourceResponseBodyData) *CreateDatasourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateDatasourceResponseBody) SetMessage(v string) *CreateDatasourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDatasourceResponseBody) SetRequestId(v string) *CreateDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasourceResponseBody) SetSuccess(v bool) *CreateDatasourceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDatasourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatasourceResponseBodyData struct {
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
}

func (s CreateDatasourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDatasourceResponseBodyData) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *CreateDatasourceResponseBodyData) SetDatasourceId(v int64) *CreateDatasourceResponseBodyData {
	s.DatasourceId = &v
	return s
}

func (s *CreateDatasourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
