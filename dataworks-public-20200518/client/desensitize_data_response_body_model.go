// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDesensitizeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesensitizeData(v string) *DesensitizeDataResponseBody
	GetDesensitizeData() *string
	SetRequestId(v string) *DesensitizeDataResponseBody
	GetRequestId() *string
}

type DesensitizeDataResponseBody struct {
	// The data returned after masking.
	//
	// example:
	//
	// {     "desensData": "**198807120510**"   }
	DesensitizeData *string `json:"DesensitizeData,omitempty" xml:"DesensitizeData,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BCA321BA-46C5-40E9-8A08-6******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DesensitizeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DesensitizeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DesensitizeDataResponseBody) GetDesensitizeData() *string {
	return s.DesensitizeData
}

func (s *DesensitizeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DesensitizeDataResponseBody) SetDesensitizeData(v string) *DesensitizeDataResponseBody {
	s.DesensitizeData = &v
	return s
}

func (s *DesensitizeDataResponseBody) SetRequestId(v string) *DesensitizeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DesensitizeDataResponseBody) Validate() error {
	return dara.Validate(s)
}
