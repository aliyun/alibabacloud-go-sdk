// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopApplicationRequest
	GetAppId() *string
}

type StopApplicationRequest struct {
	// The returned message.
	//
	// 	- **success*	- is returned when the request succeeds.
	//
	// 	- An error code is returned when the request fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s StopApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s StopApplicationRequest) GoString() string {
	return s.String()
}

func (s *StopApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopApplicationRequest) SetAppId(v string) *StopApplicationRequest {
	s.AppId = &v
	return s
}

func (s *StopApplicationRequest) Validate() error {
	return dara.Validate(s)
}
