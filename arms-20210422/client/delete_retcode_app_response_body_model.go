// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRetcodeAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteRetcodeAppResponseBody
	GetData() *string
	SetRequestId(v string) *DeleteRetcodeAppResponseBody
	GetRequestId() *string
}

type DeleteRetcodeAppResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRetcodeAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRetcodeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteRetcodeAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRetcodeAppResponseBody) SetData(v string) *DeleteRetcodeAppResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) SetRequestId(v string) *DeleteRetcodeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) Validate() error {
	return dara.Validate(s)
}
