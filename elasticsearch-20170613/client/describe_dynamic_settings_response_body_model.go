// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDynamicSettingsResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeDynamicSettingsResponseBody
	GetResult() *string
}

type DescribeDynamicSettingsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeDynamicSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDynamicSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDynamicSettingsResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeDynamicSettingsResponseBody) SetRequestId(v string) *DescribeDynamicSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDynamicSettingsResponseBody) SetResult(v string) *DescribeDynamicSettingsResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeDynamicSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
