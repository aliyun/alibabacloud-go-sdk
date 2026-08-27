// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTenantDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTenantDirectoryResponseBody
	GetCode() *string
	SetDeleteMode(v string) *DeleteTenantDirectoryResponseBody
	GetDeleteMode() *string
	SetDirectoryId(v string) *DeleteTenantDirectoryResponseBody
	GetDirectoryId() *string
	SetMessage(v string) *DeleteTenantDirectoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTenantDirectoryResponseBody
	GetRequestId() *string
}

type DeleteTenantDirectoryResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The deletion mode that is actually used.
	//
	// example:
	//
	// reject
	DeleteMode *string `json:"deleteMode,omitempty" xml:"deleteMode,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTenantDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTenantDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTenantDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTenantDirectoryResponseBody) GetDeleteMode() *string {
	return s.DeleteMode
}

func (s *DeleteTenantDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteTenantDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTenantDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTenantDirectoryResponseBody) SetCode(v string) *DeleteTenantDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTenantDirectoryResponseBody) SetDeleteMode(v string) *DeleteTenantDirectoryResponseBody {
	s.DeleteMode = &v
	return s
}

func (s *DeleteTenantDirectoryResponseBody) SetDirectoryId(v string) *DeleteTenantDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *DeleteTenantDirectoryResponseBody) SetMessage(v string) *DeleteTenantDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTenantDirectoryResponseBody) SetRequestId(v string) *DeleteTenantDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTenantDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
