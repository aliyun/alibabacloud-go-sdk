// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVServerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachVServerGroupsResponseBody
	GetRequestId() *string
}

type AttachVServerGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachVServerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachVServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachVServerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachVServerGroupsResponseBody) SetRequestId(v string) *AttachVServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachVServerGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
