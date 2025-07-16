// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlignStoragePrimaryAzoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AlignStoragePrimaryAzoneResponseBody
	GetMessage() *string
	SetRequestId(v string) *AlignStoragePrimaryAzoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AlignStoragePrimaryAzoneResponseBody
	GetSuccess() *bool
}

type AlignStoragePrimaryAzoneResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AlignStoragePrimaryAzoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AlignStoragePrimaryAzoneResponseBody) GoString() string {
	return s.String()
}

func (s *AlignStoragePrimaryAzoneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AlignStoragePrimaryAzoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AlignStoragePrimaryAzoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AlignStoragePrimaryAzoneResponseBody) SetMessage(v string) *AlignStoragePrimaryAzoneResponseBody {
	s.Message = &v
	return s
}

func (s *AlignStoragePrimaryAzoneResponseBody) SetRequestId(v string) *AlignStoragePrimaryAzoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AlignStoragePrimaryAzoneResponseBody) SetSuccess(v bool) *AlignStoragePrimaryAzoneResponseBody {
	s.Success = &v
	return s
}

func (s *AlignStoragePrimaryAzoneResponseBody) Validate() error {
	return dara.Validate(s)
}
