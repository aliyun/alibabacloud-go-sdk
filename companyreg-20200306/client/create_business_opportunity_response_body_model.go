// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBusinessOpportunityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateBusinessOpportunityResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateBusinessOpportunityResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateBusinessOpportunityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBusinessOpportunityResponseBody
	GetSuccess() *bool
}

type CreateBusinessOpportunityResponseBody struct {
	// example:
	//
	// NoPermission
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 6A603AA0-73BA-52B3-AC7D-0F846ECF7A9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBusinessOpportunityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBusinessOpportunityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBusinessOpportunityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateBusinessOpportunityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateBusinessOpportunityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBusinessOpportunityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBusinessOpportunityResponseBody) SetErrorCode(v string) *CreateBusinessOpportunityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateBusinessOpportunityResponseBody) SetErrorMessage(v string) *CreateBusinessOpportunityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateBusinessOpportunityResponseBody) SetRequestId(v string) *CreateBusinessOpportunityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBusinessOpportunityResponseBody) SetSuccess(v bool) *CreateBusinessOpportunityResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBusinessOpportunityResponseBody) Validate() error {
	return dara.Validate(s)
}
