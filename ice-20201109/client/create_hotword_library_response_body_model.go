// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotwordLibraryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHotwordLibraryId(v string) *CreateHotwordLibraryResponseBody
	GetHotwordLibraryId() *string
	SetRequestId(v string) *CreateHotwordLibraryResponseBody
	GetRequestId() *string
}

type CreateHotwordLibraryResponseBody struct {
	// The ID of the hotword library.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	HotwordLibraryId *string `json:"HotwordLibraryId,omitempty" xml:"HotwordLibraryId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 13cbb83e-043c-4728-ac35-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHotwordLibraryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHotwordLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHotwordLibraryResponseBody) GetHotwordLibraryId() *string {
	return s.HotwordLibraryId
}

func (s *CreateHotwordLibraryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHotwordLibraryResponseBody) SetHotwordLibraryId(v string) *CreateHotwordLibraryResponseBody {
	s.HotwordLibraryId = &v
	return s
}

func (s *CreateHotwordLibraryResponseBody) SetRequestId(v string) *CreateHotwordLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHotwordLibraryResponseBody) Validate() error {
	return dara.Validate(s)
}
