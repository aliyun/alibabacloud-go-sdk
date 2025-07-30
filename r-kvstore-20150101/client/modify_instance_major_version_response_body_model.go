// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMajorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceMajorVersionResponseBody
	GetRequestId() *string
}

type ModifyInstanceMajorVersionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AA587FB2-2593-4DFE-BE13-2494C2DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMajorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMajorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMajorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceMajorVersionResponseBody) SetRequestId(v string) *ModifyInstanceMajorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceMajorVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
