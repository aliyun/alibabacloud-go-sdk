// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLCollectorRetentionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *DescribeSQLCollectorRetentionResponseBody
	GetConfigValue() *string
	SetRequestId(v string) *DescribeSQLCollectorRetentionResponseBody
	GetRequestId() *string
}

type DescribeSQLCollectorRetentionResponseBody struct {
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLCollectorRetentionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLCollectorRetentionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorRetentionResponseBody) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeSQLCollectorRetentionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLCollectorRetentionResponseBody) SetConfigValue(v string) *DescribeSQLCollectorRetentionResponseBody {
	s.ConfigValue = &v
	return s
}

func (s *DescribeSQLCollectorRetentionResponseBody) SetRequestId(v string) *DescribeSQLCollectorRetentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLCollectorRetentionResponseBody) Validate() error {
	return dara.Validate(s)
}
