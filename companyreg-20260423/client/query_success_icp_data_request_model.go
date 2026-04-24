// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySuccessIcpDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaller(v string) *QuerySuccessIcpDataRequest
	GetCaller() *string
}

type QuerySuccessIcpDataRequest struct {
	// example:
	//
	// skill
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s QuerySuccessIcpDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySuccessIcpDataRequest) GoString() string {
	return s.String()
}

func (s *QuerySuccessIcpDataRequest) GetCaller() *string {
	return s.Caller
}

func (s *QuerySuccessIcpDataRequest) SetCaller(v string) *QuerySuccessIcpDataRequest {
	s.Caller = &v
	return s
}

func (s *QuerySuccessIcpDataRequest) Validate() error {
	return dara.Validate(s)
}
