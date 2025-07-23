// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppFailBackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AppFailBackResponseBody
	GetCode() *string
	SetData(v int32) *AppFailBackResponseBody
	GetData() *int32
	SetMessage(v string) *AppFailBackResponseBody
	GetMessage() *string
	SetRequestId(v string) *AppFailBackResponseBody
	GetRequestId() *string
}

type AppFailBackResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The disaster recovery switchback task ID.
	//
	// example:
	//
	// 3309
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// OKITHEVRQCN6ULQG
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7036DDBE-0ABA-52D7-A39D-75E511970F07
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppFailBackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AppFailBackResponseBody) GoString() string {
	return s.String()
}

func (s *AppFailBackResponseBody) GetCode() *string {
	return s.Code
}

func (s *AppFailBackResponseBody) GetData() *int32 {
	return s.Data
}

func (s *AppFailBackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AppFailBackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AppFailBackResponseBody) SetCode(v string) *AppFailBackResponseBody {
	s.Code = &v
	return s
}

func (s *AppFailBackResponseBody) SetData(v int32) *AppFailBackResponseBody {
	s.Data = &v
	return s
}

func (s *AppFailBackResponseBody) SetMessage(v string) *AppFailBackResponseBody {
	s.Message = &v
	return s
}

func (s *AppFailBackResponseBody) SetRequestId(v string) *AppFailBackResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppFailBackResponseBody) Validate() error {
	return dara.Validate(s)
}
