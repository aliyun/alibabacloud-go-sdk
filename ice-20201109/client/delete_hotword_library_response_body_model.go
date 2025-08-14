// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotwordLibraryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHotwordLibraryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHotwordLibraryResponseBody
	GetSuccess() *bool
}

type DeleteHotwordLibraryResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ****83B7-7F87-4792-BFE9-63CD2137****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHotwordLibraryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotwordLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotwordLibraryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHotwordLibraryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHotwordLibraryResponseBody) SetRequestId(v string) *DeleteHotwordLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotwordLibraryResponseBody) SetSuccess(v bool) *DeleteHotwordLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHotwordLibraryResponseBody) Validate() error {
	return dara.Validate(s)
}
