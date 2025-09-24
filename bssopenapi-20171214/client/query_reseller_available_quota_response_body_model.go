// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResellerAvailableQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryResellerAvailableQuotaResponseBody
	GetCode() *string
	SetData(v string) *QueryResellerAvailableQuotaResponseBody
	GetData() *string
	SetMessage(v string) *QueryResellerAvailableQuotaResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryResellerAvailableQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryResellerAvailableQuotaResponseBody
	GetSuccess() *bool
}

type QueryResellerAvailableQuotaResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 300
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryResellerAvailableQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryResellerAvailableQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResellerAvailableQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryResellerAvailableQuotaResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryResellerAvailableQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryResellerAvailableQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryResellerAvailableQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryResellerAvailableQuotaResponseBody) SetCode(v string) *QueryResellerAvailableQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponseBody) SetData(v string) *QueryResellerAvailableQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponseBody) SetMessage(v string) *QueryResellerAvailableQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponseBody) SetRequestId(v string) *QueryResellerAvailableQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponseBody) SetSuccess(v bool) *QueryResellerAvailableQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
