// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceRefreshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeInstanceRefreshResponseBody
	GetRequestId() *string
}

type ResumeInstanceRefreshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeInstanceRefreshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRefreshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeInstanceRefreshResponseBody) SetRequestId(v string) *ResumeInstanceRefreshResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeInstanceRefreshResponseBody) Validate() error {
	return dara.Validate(s)
}
