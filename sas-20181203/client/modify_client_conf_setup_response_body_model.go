// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientConfSetupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClientConfSetupResponseBody
	GetRequestId() *string
}

type ModifyClientConfSetupResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 60E24426-B910-5D7F-8B8B-3BCDC3FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClientConfSetupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientConfSetupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClientConfSetupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClientConfSetupResponseBody) SetRequestId(v string) *ModifyClientConfSetupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClientConfSetupResponseBody) Validate() error {
	return dara.Validate(s)
}
