// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequest(v map[string]interface{}) *ListCalendarsRequest
	GetRequest() map[string]interface{}
}

type ListCalendarsRequest struct {
	// example:
	//
	// {}
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s ListCalendarsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsRequest) GoString() string {
	return s.String()
}

func (s *ListCalendarsRequest) GetRequest() map[string]interface{} {
	return s.Request
}

func (s *ListCalendarsRequest) SetRequest(v map[string]interface{}) *ListCalendarsRequest {
	s.Request = v
	return s
}

func (s *ListCalendarsRequest) Validate() error {
	return dara.Validate(s)
}
