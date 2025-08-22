// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListCalendarNamesRequest
	GetClusterId() *string
}

type ListCalendarNamesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListCalendarNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarNamesRequest) GoString() string {
	return s.String()
}

func (s *ListCalendarNamesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListCalendarNamesRequest) SetClusterId(v string) *ListCalendarNamesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCalendarNamesRequest) Validate() error {
	return dara.Validate(s)
}
