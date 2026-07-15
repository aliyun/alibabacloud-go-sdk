// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *ImportMediaResponseBody
	GetMediaId() *string
	SetRequestId(v string) *ImportMediaResponseBody
	GetRequestId() *string
}

type ImportMediaResponseBody struct {
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportMediaResponseBody) GoString() string {
	return s.String()
}

func (s *ImportMediaResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *ImportMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportMediaResponseBody) SetMediaId(v string) *ImportMediaResponseBody {
	s.MediaId = &v
	return s
}

func (s *ImportMediaResponseBody) SetRequestId(v string) *ImportMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
