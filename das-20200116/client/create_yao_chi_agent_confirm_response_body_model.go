// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYaoChiAgentConfirmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApproved(v bool) *CreateYaoChiAgentConfirmResponseBody
	GetApproved() *bool
	SetConfirmationId(v string) *CreateYaoChiAgentConfirmResponseBody
	GetConfirmationId() *string
	SetRequestId(v string) *CreateYaoChiAgentConfirmResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateYaoChiAgentConfirmResponseBody
	GetStatus() *string
	SetToolArgsReceived(v bool) *CreateYaoChiAgentConfirmResponseBody
	GetToolArgsReceived() *bool
}

type CreateYaoChiAgentConfirmResponseBody struct {
	// The confirmation decision submitted in this request, which is the same as the Approved request parameter. A value of true indicates that the execution is approved, and false indicates that the execution is rejected. This does not indicate whether the tool operation is executed successfully.
	//
	// example:
	//
	// true
	Approved *bool `json:"Approved,omitempty" xml:"Approved,omitempty"`
	// The identifier of the operation for which the confirmation is submitted, which is the same as the ConfirmationId request parameter.
	//
	// example:
	//
	// 123e4567-e89b-42d3-a456-426614174000
	ConfirmationId *string `json:"ConfirmationId,omitempty" xml:"ConfirmationId,omitempty"`
	// The unique identifier of the request, which is used for troubleshooting and log tracing.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B23456
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status string. The example value success refers to the outer success status of the downstream confirmation callback response and is provided only for format demonstration. It does not mean that the Status field of the current operation actually returns this value. In the current version, this field may be empty or not returned. Do not rely on this field to determine whether the execution is approved or whether the tool is executed successfully.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether custom tool parameters are received. A value of true indicates that non-null parameters are provided, and false indicates that no parameters are provided. This does not indicate that the parameters are applied or that the tool is executed successfully. In the current version, omit the ToolArgs request parameter, and this field normally returns false.
	//
	// example:
	//
	// false
	ToolArgsReceived *bool `json:"ToolArgsReceived,omitempty" xml:"ToolArgsReceived,omitempty"`
}

func (s CreateYaoChiAgentConfirmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateYaoChiAgentConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateYaoChiAgentConfirmResponseBody) GetApproved() *bool {
	return s.Approved
}

func (s *CreateYaoChiAgentConfirmResponseBody) GetConfirmationId() *string {
	return s.ConfirmationId
}

func (s *CreateYaoChiAgentConfirmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateYaoChiAgentConfirmResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateYaoChiAgentConfirmResponseBody) GetToolArgsReceived() *bool {
	return s.ToolArgsReceived
}

func (s *CreateYaoChiAgentConfirmResponseBody) SetApproved(v bool) *CreateYaoChiAgentConfirmResponseBody {
	s.Approved = &v
	return s
}

func (s *CreateYaoChiAgentConfirmResponseBody) SetConfirmationId(v string) *CreateYaoChiAgentConfirmResponseBody {
	s.ConfirmationId = &v
	return s
}

func (s *CreateYaoChiAgentConfirmResponseBody) SetRequestId(v string) *CreateYaoChiAgentConfirmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateYaoChiAgentConfirmResponseBody) SetStatus(v string) *CreateYaoChiAgentConfirmResponseBody {
	s.Status = &v
	return s
}

func (s *CreateYaoChiAgentConfirmResponseBody) SetToolArgsReceived(v bool) *CreateYaoChiAgentConfirmResponseBody {
	s.ToolArgsReceived = &v
	return s
}

func (s *CreateYaoChiAgentConfirmResponseBody) Validate() error {
	return dara.Validate(s)
}
