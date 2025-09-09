// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceStoreInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMax(v int64) *GetInstanceStoreInfoResponseBody
	GetMax() *int64
	SetRequestId(v string) *GetInstanceStoreInfoResponseBody
	GetRequestId() *string
	SetUsage(v int64) *GetInstanceStoreInfoResponseBody
	GetUsage() *int64
}

type GetInstanceStoreInfoResponseBody struct {
	// example:
	//
	// 102400
	Max *int64 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 5EAB922E-F476-5DFA-9290-313C608E724B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1024
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s GetInstanceStoreInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceStoreInfoResponseBody) GetMax() *int64 {
	return s.Max
}

func (s *GetInstanceStoreInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceStoreInfoResponseBody) GetUsage() *int64 {
	return s.Usage
}

func (s *GetInstanceStoreInfoResponseBody) SetMax(v int64) *GetInstanceStoreInfoResponseBody {
	s.Max = &v
	return s
}

func (s *GetInstanceStoreInfoResponseBody) SetRequestId(v string) *GetInstanceStoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceStoreInfoResponseBody) SetUsage(v int64) *GetInstanceStoreInfoResponseBody {
	s.Usage = &v
	return s
}

func (s *GetInstanceStoreInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
