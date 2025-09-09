// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlashbackRecycleBinTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *FlashbackRecycleBinTableResponseBody
	GetData() *bool
	SetRequestId(v string) *FlashbackRecycleBinTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlashbackRecycleBinTableResponseBody
	GetSuccess() *bool
}

type FlashbackRecycleBinTableResponseBody struct {
	// Indicates whether the deleted logical table is restored.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 463A5F0F-12AD-4544-A902-B2B983******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FlashbackRecycleBinTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlashbackRecycleBinTableResponseBody) GoString() string {
	return s.String()
}

func (s *FlashbackRecycleBinTableResponseBody) GetData() *bool {
	return s.Data
}

func (s *FlashbackRecycleBinTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlashbackRecycleBinTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlashbackRecycleBinTableResponseBody) SetData(v bool) *FlashbackRecycleBinTableResponseBody {
	s.Data = &v
	return s
}

func (s *FlashbackRecycleBinTableResponseBody) SetRequestId(v string) *FlashbackRecycleBinTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlashbackRecycleBinTableResponseBody) SetSuccess(v bool) *FlashbackRecycleBinTableResponseBody {
	s.Success = &v
	return s
}

func (s *FlashbackRecycleBinTableResponseBody) Validate() error {
	return dara.Validate(s)
}
