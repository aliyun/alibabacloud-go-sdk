// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v interface{}) *ModifyScheduledTaskRequest
	GetBody() interface{}
}

type ModifyScheduledTaskRequest struct {
	// The request parameters.
	//
	// example:
	//
	// {
	//
	//   "cron": "0 0 	- 	- 1,2,3,4,5,6,7",
	//
	//   "enabled": true,
	//
	//   "filter": {
	//
	//     "field": "cat_id",
	//
	//     "days": 30,
	//
	//     "unit": "s"
	//
	//   }
	//
	// }
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequest) GetBody() interface{} {
	return s.Body
}

func (s *ModifyScheduledTaskRequest) SetBody(v interface{}) *ModifyScheduledTaskRequest {
	s.Body = v
	return s
}

func (s *ModifyScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
