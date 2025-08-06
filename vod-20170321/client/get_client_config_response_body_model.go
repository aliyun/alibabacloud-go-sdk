// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetClientConfigResponseBody
	GetData() *string
	SetRequestId(v string) *GetClientConfigResponseBody
	GetRequestId() *string
}

type GetClientConfigResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *GetClientConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientConfigResponseBody) SetData(v string) *GetClientConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetClientConfigResponseBody) SetRequestId(v string) *GetClientConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
