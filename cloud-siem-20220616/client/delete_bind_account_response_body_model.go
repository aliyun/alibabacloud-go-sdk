// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBindAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteBindAccountResponseBodyData) *DeleteBindAccountResponseBody
	GetData() *DeleteBindAccountResponseBodyData
	SetRequestId(v string) *DeleteBindAccountResponseBody
	GetRequestId() *string
}

type DeleteBindAccountResponseBody struct {
	// The data returned.
	Data *DeleteBindAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBindAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBindAccountResponseBody) GetData() *DeleteBindAccountResponseBodyData {
	return s.Data
}

func (s *DeleteBindAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBindAccountResponseBody) SetData(v *DeleteBindAccountResponseBodyData) *DeleteBindAccountResponseBody {
	s.Data = v
	return s
}

func (s *DeleteBindAccountResponseBody) SetRequestId(v string) *DeleteBindAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBindAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteBindAccountResponseBodyData struct {
	// The number of cloud accounts that are removed. The value 1 indicates that cloud account is removed, and a value less than or equal to 0 indicates that the cloud account failed to be removed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DeleteBindAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteBindAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteBindAccountResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *DeleteBindAccountResponseBodyData) SetCount(v int32) *DeleteBindAccountResponseBodyData {
	s.Count = &v
	return s
}

func (s *DeleteBindAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
