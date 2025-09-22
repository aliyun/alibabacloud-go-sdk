// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAllActivityInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *QueryAllActivityInfoRequest
	GetActivityId() *string
}

type QueryAllActivityInfoRequest struct {
	// This parameter is required.
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
}

func (s QueryAllActivityInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAllActivityInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAllActivityInfoRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *QueryAllActivityInfoRequest) SetActivityId(v string) *QueryAllActivityInfoRequest {
	s.ActivityId = &v
	return s
}

func (s *QueryAllActivityInfoRequest) Validate() error {
	return dara.Validate(s)
}
