// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashscopeAsyncTaskFinishEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]interface{}) *DashscopeAsyncTaskFinishEventRequest
	GetBody() map[string]interface{}
}

type DashscopeAsyncTaskFinishEventRequest struct {
	// example:
	//
	// {}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DashscopeAsyncTaskFinishEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DashscopeAsyncTaskFinishEventRequest) GoString() string {
	return s.String()
}

func (s *DashscopeAsyncTaskFinishEventRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *DashscopeAsyncTaskFinishEventRequest) SetBody(v map[string]interface{}) *DashscopeAsyncTaskFinishEventRequest {
	s.Body = v
	return s
}

func (s *DashscopeAsyncTaskFinishEventRequest) Validate() error {
	return dara.Validate(s)
}
