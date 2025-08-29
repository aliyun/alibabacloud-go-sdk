// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCalculationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetCalculationJobRequest
	GetInstanceId() *string
}

type GetCalculationJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCalculationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCalculationJobRequest) GoString() string {
	return s.String()
}

func (s *GetCalculationJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCalculationJobRequest) SetInstanceId(v string) *GetCalculationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCalculationJobRequest) Validate() error {
	return dara.Validate(s)
}
