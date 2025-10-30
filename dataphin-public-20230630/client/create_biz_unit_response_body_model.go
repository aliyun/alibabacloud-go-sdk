// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBizUnitResponseBody
	GetCode() *string
	SetCreateResult(v *CreateBizUnitResponseBodyCreateResult) *CreateBizUnitResponseBody
	GetCreateResult() *CreateBizUnitResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateBizUnitResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBizUnitResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBizUnitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBizUnitResponseBody
	GetSuccess() *bool
}

type CreateBizUnitResponseBody struct {
	// example:
	//
	// OK
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateResult *CreateBizUnitResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBizUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBizUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBizUnitResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBizUnitResponseBody) GetCreateResult() *CreateBizUnitResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateBizUnitResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBizUnitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBizUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBizUnitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBizUnitResponseBody) SetCode(v string) *CreateBizUnitResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBizUnitResponseBody) SetCreateResult(v *CreateBizUnitResponseBodyCreateResult) *CreateBizUnitResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateBizUnitResponseBody) SetHttpStatusCode(v int32) *CreateBizUnitResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBizUnitResponseBody) SetMessage(v string) *CreateBizUnitResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBizUnitResponseBody) SetRequestId(v string) *CreateBizUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBizUnitResponseBody) SetSuccess(v bool) *CreateBizUnitResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBizUnitResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBizUnitResponseBodyCreateResult struct {
	// example:
	//
	// 545844456
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
}

func (s CreateBizUnitResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateBizUnitResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateBizUnitResponseBodyCreateResult) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *CreateBizUnitResponseBodyCreateResult) SetBizUnitId(v int64) *CreateBizUnitResponseBodyCreateResult {
	s.BizUnitId = &v
	return s
}

func (s *CreateBizUnitResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
