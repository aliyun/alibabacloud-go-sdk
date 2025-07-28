// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptFulfillmentDecisionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AcceptFulfillmentDecisionResponseBody
	GetData() *bool
	SetErrorCode(v string) *AcceptFulfillmentDecisionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AcceptFulfillmentDecisionResponseBody
	GetErrorMessage() *string
	SetSuccess(v bool) *AcceptFulfillmentDecisionResponseBody
	GetSuccess() *bool
}

type AcceptFulfillmentDecisionResponseBody struct {
	Data         *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AcceptFulfillmentDecisionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptFulfillmentDecisionResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptFulfillmentDecisionResponseBody) GetData() *bool {
	return s.Data
}

func (s *AcceptFulfillmentDecisionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AcceptFulfillmentDecisionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AcceptFulfillmentDecisionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AcceptFulfillmentDecisionResponseBody) SetData(v bool) *AcceptFulfillmentDecisionResponseBody {
	s.Data = &v
	return s
}

func (s *AcceptFulfillmentDecisionResponseBody) SetErrorCode(v string) *AcceptFulfillmentDecisionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AcceptFulfillmentDecisionResponseBody) SetErrorMessage(v string) *AcceptFulfillmentDecisionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AcceptFulfillmentDecisionResponseBody) SetSuccess(v bool) *AcceptFulfillmentDecisionResponseBody {
	s.Success = &v
	return s
}

func (s *AcceptFulfillmentDecisionResponseBody) Validate() error {
	return dara.Validate(s)
}
