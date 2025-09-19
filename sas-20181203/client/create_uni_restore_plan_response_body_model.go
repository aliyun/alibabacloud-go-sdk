// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUniRestorePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUniRestorePlanResponseBody
	GetRequestId() *string
}

type CreateUniRestorePlanResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F5CF78A7-30AA-59DB-847F-13EE3AE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUniRestorePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUniRestorePlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUniRestorePlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUniRestorePlanResponseBody) SetRequestId(v string) *CreateUniRestorePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUniRestorePlanResponseBody) Validate() error {
	return dara.Validate(s)
}
