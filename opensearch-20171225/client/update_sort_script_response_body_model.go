// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSortScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSortScriptResponseBody
	GetRequestId() *string
}

type UpdateSortScriptResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9F165784-5507-5342-ABF3-545293F9808A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSortScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSortScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSortScriptResponseBody) SetRequestId(v string) *UpdateSortScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSortScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
