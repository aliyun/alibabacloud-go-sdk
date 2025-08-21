// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteApplicationRequest
	GetAppId() *string
	SetTimeout(v int32) *DeleteApplicationRequest
	GetTimeout() *int32
}

type DeleteApplicationRequest struct {
	// The ID of the application. To obtain the application ID, call the ListApplication operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// d0639abf-789a-4527-b420-031d2cd9ad9b
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The timeout period for the asynchronous release. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 1800
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeleteApplicationRequest) SetAppId(v string) *DeleteApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeleteApplicationRequest) SetTimeout(v int32) *DeleteApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *DeleteApplicationRequest) Validate() error {
	return dara.Validate(s)
}
