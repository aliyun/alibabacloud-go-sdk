// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProcessInstanceId(v string) *GetApplicationContentsRequest
	GetProcessInstanceId() *string
}

type GetApplicationContentsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 332066440109224007
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
}

func (s GetApplicationContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationContentsRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationContentsRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetApplicationContentsRequest) SetProcessInstanceId(v string) *GetApplicationContentsRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetApplicationContentsRequest) Validate() error {
	return dara.Validate(s)
}
