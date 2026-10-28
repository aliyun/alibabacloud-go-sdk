// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReason(v string) *CancelRunRequest
	GetReason() *string
}

type CancelRunRequest struct {
	// The reason for the cancellation. The value must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// The user manually canceled this task
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s CancelRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelRunRequest) GoString() string {
	return s.String()
}

func (s *CancelRunRequest) GetReason() *string {
	return s.Reason
}

func (s *CancelRunRequest) SetReason(v string) *CancelRunRequest {
	s.Reason = &v
	return s
}

func (s *CancelRunRequest) Validate() error {
	return dara.Validate(s)
}
