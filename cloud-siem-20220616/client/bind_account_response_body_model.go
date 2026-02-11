// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BindAccountResponseBodyData) *BindAccountResponseBody
	GetData() *BindAccountResponseBodyData
	SetRequestId(v string) *BindAccountResponseBody
	GetRequestId() *string
}

type BindAccountResponseBody struct {
	// The data returned.
	Data *BindAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *BindAccountResponseBody) GetData() *BindAccountResponseBodyData {
	return s.Data
}

func (s *BindAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAccountResponseBody) SetData(v *BindAccountResponseBodyData) *BindAccountResponseBody {
	s.Data = v
	return s
}

func (s *BindAccountResponseBody) SetRequestId(v string) *BindAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAccountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAccountResponseBodyData struct {
	// The number of the cloud accounts that are added to the threat analysis feature.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s BindAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindAccountResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *BindAccountResponseBodyData) SetCount(v int32) *BindAccountResponseBodyData {
	s.Count = &v
	return s
}

func (s *BindAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
