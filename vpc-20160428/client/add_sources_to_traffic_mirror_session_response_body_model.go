// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSourcesToTrafficMirrorSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddSourcesToTrafficMirrorSessionResponseBody
	GetRequestId() *string
}

type AddSourcesToTrafficMirrorSessionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 87F70089-5B38-41A8-BAD8-0B55E2F8DC57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSourcesToTrafficMirrorSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSourcesToTrafficMirrorSessionResponseBody) GoString() string {
	return s.String()
}

func (s *AddSourcesToTrafficMirrorSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSourcesToTrafficMirrorSessionResponseBody) SetRequestId(v string) *AddSourcesToTrafficMirrorSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
