// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTensorboardSharedUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTensorboardSharedUrlResponseBody
	GetRequestId() *string
	SetTensorboardSharedUrl(v string) *GetTensorboardSharedUrlResponseBody
	GetTensorboardSharedUrl() *string
}

type GetTensorboardSharedUrlResponseBody struct {
	// The request ID which is used for troubleshooting.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The shareable link of the TensorBoard task.
	//
	// example:
	//
	// http://pai-dlc-proxy-xxx.alicyuncs.com/xxx/xxx/token/
	TensorboardSharedUrl *string `json:"TensorboardSharedUrl,omitempty" xml:"TensorboardSharedUrl,omitempty"`
}

func (s GetTensorboardSharedUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTensorboardSharedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTensorboardSharedUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTensorboardSharedUrlResponseBody) GetTensorboardSharedUrl() *string {
	return s.TensorboardSharedUrl
}

func (s *GetTensorboardSharedUrlResponseBody) SetRequestId(v string) *GetTensorboardSharedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTensorboardSharedUrlResponseBody) SetTensorboardSharedUrl(v string) *GetTensorboardSharedUrlResponseBody {
	s.TensorboardSharedUrl = &v
	return s
}

func (s *GetTensorboardSharedUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
