// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCnameFlatteningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlattenMode(v string) *GetCnameFlatteningResponseBody
	GetFlattenMode() *string
	SetRequestId(v string) *GetCnameFlatteningResponseBody
	GetRequestId() *string
}

type GetCnameFlatteningResponseBody struct {
	FlattenMode *string `json:"FlattenMode,omitempty" xml:"FlattenMode,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCnameFlatteningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCnameFlatteningResponseBody) GoString() string {
	return s.String()
}

func (s *GetCnameFlatteningResponseBody) GetFlattenMode() *string {
	return s.FlattenMode
}

func (s *GetCnameFlatteningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCnameFlatteningResponseBody) SetFlattenMode(v string) *GetCnameFlatteningResponseBody {
	s.FlattenMode = &v
	return s
}

func (s *GetCnameFlatteningResponseBody) SetRequestId(v string) *GetCnameFlatteningResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCnameFlatteningResponseBody) Validate() error {
	return dara.Validate(s)
}
