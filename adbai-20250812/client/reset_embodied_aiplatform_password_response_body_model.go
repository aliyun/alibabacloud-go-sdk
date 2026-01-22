// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetEmbodiedAIPlatformPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetEmbodiedAIPlatformPasswordResponseBody
	GetRequestId() *string
}

type ResetEmbodiedAIPlatformPasswordResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetEmbodiedAIPlatformPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetEmbodiedAIPlatformPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetEmbodiedAIPlatformPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetEmbodiedAIPlatformPasswordResponseBody) SetRequestId(v string) *ResetEmbodiedAIPlatformPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetEmbodiedAIPlatformPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
