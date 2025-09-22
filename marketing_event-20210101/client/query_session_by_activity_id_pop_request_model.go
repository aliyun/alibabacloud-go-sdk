// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionByActivityIdPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *QuerySessionByActivityIdPopRequest
	GetActivityId() *int64
}

type QuerySessionByActivityIdPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
}

func (s QuerySessionByActivityIdPopRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionByActivityIdPopRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionByActivityIdPopRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *QuerySessionByActivityIdPopRequest) SetActivityId(v int64) *QuerySessionByActivityIdPopRequest {
	s.ActivityId = &v
	return s
}

func (s *QuerySessionByActivityIdPopRequest) Validate() error {
	return dara.Validate(s)
}
