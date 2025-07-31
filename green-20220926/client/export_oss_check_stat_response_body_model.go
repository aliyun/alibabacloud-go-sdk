// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportOssCheckStatResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *ExportOssCheckStatResponseBody
  GetData() *string 
  SetRequestId(v string) *ExportOssCheckStatResponseBody
  GetRequestId() *string 
}

type ExportOssCheckStatResponseBody struct {
  // example:
  // 
  // https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/console_data/production/scanResult/osscheck/ossCheckStat_aliUf5B3lJfOkLpqozLIn94Uy-1XxKyX.xlsx
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportOssCheckStatResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportOssCheckStatResponseBody) GoString() string {
  return s.String()
}

func (s *ExportOssCheckStatResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportOssCheckStatResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportOssCheckStatResponseBody) SetData(v string) *ExportOssCheckStatResponseBody {
  s.Data = &v
  return s
}

func (s *ExportOssCheckStatResponseBody) SetRequestId(v string) *ExportOssCheckStatResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportOssCheckStatResponseBody) Validate() error {
  return dara.Validate(s)
}

