// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNewestInnerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequest(v map[string]interface{}) *GetNewestInnerGroupsRequest
	GetRequest() map[string]interface{}
}

type GetNewestInnerGroupsRequest struct {
	// example:
	//
	// {}
	Request map[string]interface{} `json:"Request,omitempty" xml:"Request,omitempty"`
}

func (s GetNewestInnerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNewestInnerGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsRequest) GetRequest() map[string]interface{} {
	return s.Request
}

func (s *GetNewestInnerGroupsRequest) SetRequest(v map[string]interface{}) *GetNewestInnerGroupsRequest {
	s.Request = v
	return s
}

func (s *GetNewestInnerGroupsRequest) Validate() error {
	return dara.Validate(s)
}
