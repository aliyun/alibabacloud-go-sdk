// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectEventCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int32) *GetFileProtectEventCountResponseBody
	GetData() *int32
	SetRequestId(v string) *GetFileProtectEventCountResponseBody
	GetRequestId() *string
}

type GetFileProtectEventCountResponseBody struct {
	// The data returned if the request is successful.
	//
	// example:
	//
	// 16
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B28EC81-2FA7-5097-80D9-0DBE1A3DBD59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileProtectEventCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectEventCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileProtectEventCountResponseBody) GetData() *int32 {
	return s.Data
}

func (s *GetFileProtectEventCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileProtectEventCountResponseBody) SetData(v int32) *GetFileProtectEventCountResponseBody {
	s.Data = &v
	return s
}

func (s *GetFileProtectEventCountResponseBody) SetRequestId(v string) *GetFileProtectEventCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileProtectEventCountResponseBody) Validate() error {
	return dara.Validate(s)
}
