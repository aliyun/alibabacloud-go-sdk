// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagValue(v string) *DeleteQuotaRequest
	GetTagValue() *string
}

type DeleteQuotaRequest struct {
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s DeleteQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaRequest) GoString() string {
	return s.String()
}

func (s *DeleteQuotaRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *DeleteQuotaRequest) SetTagValue(v string) *DeleteQuotaRequest {
	s.TagValue = &v
	return s
}

func (s *DeleteQuotaRequest) Validate() error {
	return dara.Validate(s)
}
