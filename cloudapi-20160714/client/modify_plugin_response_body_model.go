// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPluginResponseBody
	GetRequestId() *string
}

type ModifyPluginResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CE5722A6-AE78-4741-A9B0-6C817D360510
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPluginResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPluginResponseBody) SetRequestId(v string) *ModifyPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
