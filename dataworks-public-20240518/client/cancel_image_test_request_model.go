// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelImageTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CancelImageTestRequest
	GetId() *string
	SetProcessId(v string) *CancelImageTestRequest
	GetProcessId() *string
}

type CancelImageTestRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the image test execution to cancel.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s CancelImageTestRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelImageTestRequest) GoString() string {
	return s.String()
}

func (s *CancelImageTestRequest) GetId() *string {
	return s.Id
}

func (s *CancelImageTestRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *CancelImageTestRequest) SetId(v string) *CancelImageTestRequest {
	s.Id = &v
	return s
}

func (s *CancelImageTestRequest) SetProcessId(v string) *CancelImageTestRequest {
	s.ProcessId = &v
	return s
}

func (s *CancelImageTestRequest) Validate() error {
	return dara.Validate(s)
}
