// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotwordLibraryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHotwordLibraryResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateHotwordLibraryResponseBody
	GetSuccess() *string
}

type UpdateHotwordLibraryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// *3B-0E1A-586A-AC29-742247*
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the hotword library.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHotwordLibraryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotwordLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotwordLibraryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHotwordLibraryResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateHotwordLibraryResponseBody) SetRequestId(v string) *UpdateHotwordLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotwordLibraryResponseBody) SetSuccess(v string) *UpdateHotwordLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHotwordLibraryResponseBody) Validate() error {
	return dara.Validate(s)
}
