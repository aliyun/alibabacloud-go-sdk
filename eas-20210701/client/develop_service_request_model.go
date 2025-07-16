// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDevelopServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExit(v string) *DevelopServiceRequest
	GetExit() *string
}

type DevelopServiceRequest struct {
	// Specifies whether to exit development mode. Valid values:
	//
	// 	- true: exits development mode.
	//
	// 	- false (default): enters development mode.
	//
	// example:
	//
	// true
	Exit *string `json:"Exit,omitempty" xml:"Exit,omitempty"`
}

func (s DevelopServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DevelopServiceRequest) GoString() string {
	return s.String()
}

func (s *DevelopServiceRequest) GetExit() *string {
	return s.Exit
}

func (s *DevelopServiceRequest) SetExit(v string) *DevelopServiceRequest {
	s.Exit = &v
	return s
}

func (s *DevelopServiceRequest) Validate() error {
	return dara.Validate(s)
}
