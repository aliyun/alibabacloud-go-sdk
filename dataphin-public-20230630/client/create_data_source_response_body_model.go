// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDataSourceResponseBody
	GetCode() *string
	SetCreateResult(v *CreateDataSourceResponseBodyCreateResult) *CreateDataSourceResponseBody
	GetCreateResult() *CreateDataSourceResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateDataSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDataSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataSourceResponseBody
	GetSuccess() *bool
}

type CreateDataSourceResponseBody struct {
	// example:
	//
	// OK
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateResult *CreateDataSourceResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataSourceResponseBody) GetCreateResult() *CreateDataSourceResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateDataSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDataSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataSourceResponseBody) SetCode(v string) *CreateDataSourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetCreateResult(v *CreateDataSourceResponseBodyCreateResult) *CreateDataSourceResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateDataSourceResponseBody) SetHttpStatusCode(v int32) *CreateDataSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetMessage(v string) *CreateDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetSuccess(v bool) *CreateDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataSourceResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataSourceResponseBodyCreateResult struct {
	// example:
	//
	// 123
	DevDataSourceId *int64 `json:"DevDataSourceId,omitempty" xml:"DevDataSourceId,omitempty"`
	// example:
	//
	// 12345
	ProdDataSourceId *int64 `json:"ProdDataSourceId,omitempty" xml:"ProdDataSourceId,omitempty"`
}

func (s CreateDataSourceResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBodyCreateResult) GetDevDataSourceId() *int64 {
	return s.DevDataSourceId
}

func (s *CreateDataSourceResponseBodyCreateResult) GetProdDataSourceId() *int64 {
	return s.ProdDataSourceId
}

func (s *CreateDataSourceResponseBodyCreateResult) SetDevDataSourceId(v int64) *CreateDataSourceResponseBodyCreateResult {
	s.DevDataSourceId = &v
	return s
}

func (s *CreateDataSourceResponseBodyCreateResult) SetProdDataSourceId(v int64) *CreateDataSourceResponseBodyCreateResult {
	s.ProdDataSourceId = &v
	return s
}

func (s *CreateDataSourceResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
