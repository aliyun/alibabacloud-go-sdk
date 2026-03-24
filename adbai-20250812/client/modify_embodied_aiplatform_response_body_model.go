// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmbodiedAIPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEmbodiedAIPlatformResponseBody
	GetRequestId() *string
}

type ModifyEmbodiedAIPlatformResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEmbodiedAIPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEmbodiedAIPlatformResponseBody) SetRequestId(v string) *ModifyEmbodiedAIPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
