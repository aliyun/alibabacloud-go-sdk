// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCollectorResponseBody
	GetRequestId() *string
	SetResult(v *CreateCollectorResponseBodyResult) *CreateCollectorResponseBody
	GetResult() *CreateCollectorResponseBodyResult
}

type CreateCollectorResponseBody struct {
	// example:
	//
	// 8466BDFB-C513-4B8D-B4E3-5AB256AB****
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateCollectorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCollectorResponseBody) GetResult() *CreateCollectorResponseBodyResult {
	return s.Result
}

func (s *CreateCollectorResponseBody) SetRequestId(v string) *CreateCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCollectorResponseBody) SetResult(v *CreateCollectorResponseBodyResult) *CreateCollectorResponseBody {
	s.Result = v
	return s
}

func (s *CreateCollectorResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCollectorResponseBodyResult struct {
	// example:
	//
	// ct-cn-4135is2tj194p****
	ResId *string `json:"resId,omitempty" xml:"resId,omitempty"`
}

func (s CreateCollectorResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCollectorResponseBodyResult) GetResId() *string {
	return s.ResId
}

func (s *CreateCollectorResponseBodyResult) SetResId(v string) *CreateCollectorResponseBodyResult {
	s.ResId = &v
	return s
}

func (s *CreateCollectorResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
