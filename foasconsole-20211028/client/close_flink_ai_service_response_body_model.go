// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseFlinkAiServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseFlinkAiServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CloseFlinkAiServiceResponseBody
	GetSuccess() *bool
}

type CloseFlinkAiServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseFlinkAiServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseFlinkAiServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CloseFlinkAiServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseFlinkAiServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloseFlinkAiServiceResponseBody) SetRequestId(v string) *CloseFlinkAiServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseFlinkAiServiceResponseBody) SetSuccess(v bool) *CloseFlinkAiServiceResponseBody {
	s.Success = &v
	return s
}

func (s *CloseFlinkAiServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
