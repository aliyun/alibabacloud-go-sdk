// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDashboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *GetDashboardResponseBody
	GetUrl() *string
}

type GetDashboardResponseBody struct {
	// The Dashboard URL
	//
	// example:
	//
	// https://dlcj1jzm1p01saqw-spark.pre-dsw-gateway-cn-hangzhou.data.aliyun.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDashboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetDashboardResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetDashboardResponseBody) SetUrl(v string) *GetDashboardResponseBody {
	s.Url = &v
	return s
}

func (s *GetDashboardResponseBody) Validate() error {
	return dara.Validate(s)
}
