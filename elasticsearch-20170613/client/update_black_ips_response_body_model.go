// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBlackIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateBlackIpsResponseBody
	GetRequestId() *string
	SetResult(v *UpdateBlackIpsResponseBodyResult) *UpdateBlackIpsResponseBody
	GetResult() *UpdateBlackIpsResponseBodyResult
}

type UpdateBlackIpsResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateBlackIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateBlackIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBlackIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBlackIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBlackIpsResponseBody) GetResult() *UpdateBlackIpsResponseBodyResult {
	return s.Result
}

func (s *UpdateBlackIpsResponseBody) SetRequestId(v string) *UpdateBlackIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBlackIpsResponseBody) SetResult(v *UpdateBlackIpsResponseBodyResult) *UpdateBlackIpsResponseBody {
	s.Result = v
	return s
}

func (s *UpdateBlackIpsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBlackIpsResponseBodyResult struct {
	EsIPBlacklist []*string `json:"esIPBlacklist,omitempty" xml:"esIPBlacklist,omitempty" type:"Repeated"`
}

func (s UpdateBlackIpsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateBlackIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateBlackIpsResponseBodyResult) GetEsIPBlacklist() []*string {
	return s.EsIPBlacklist
}

func (s *UpdateBlackIpsResponseBodyResult) SetEsIPBlacklist(v []*string) *UpdateBlackIpsResponseBodyResult {
	s.EsIPBlacklist = v
	return s
}

func (s *UpdateBlackIpsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
