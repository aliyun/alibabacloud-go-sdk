// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sAccessInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteK8sAccessInfoResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteK8sAccessInfoResponseBody
	GetRequestId() *string
}

type DeleteK8sAccessInfoResponseBody struct {
	// Indicates whether the request was successful. Valid value:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C5338DE5-5D80-51A1-B330-98300AFB80E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteK8sAccessInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sAccessInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteK8sAccessInfoResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteK8sAccessInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteK8sAccessInfoResponseBody) SetData(v bool) *DeleteK8sAccessInfoResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteK8sAccessInfoResponseBody) SetRequestId(v string) *DeleteK8sAccessInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteK8sAccessInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
