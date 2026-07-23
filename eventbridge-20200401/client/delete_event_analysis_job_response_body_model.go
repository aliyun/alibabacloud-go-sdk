// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventAnalysisJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEventAnalysisJobResponseBody
	GetRequestId() *string
}

type DeleteEventAnalysisJobResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEventAnalysisJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventAnalysisJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventAnalysisJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventAnalysisJobResponseBody) SetRequestId(v string) *DeleteEventAnalysisJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventAnalysisJobResponseBody) Validate() error {
	return dara.Validate(s)
}
