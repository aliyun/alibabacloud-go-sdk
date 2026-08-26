// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditShowAndReplaceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetJobInfo(v string) *EditShowAndReplaceResponseBody
  GetJobInfo() *string 
  SetRequestId(v string) *EditShowAndReplaceResponseBody
  GetRequestId() *string 
}

type EditShowAndReplaceResponseBody struct {
  // The task information, which contains:
  // 
  // - **vodId**: the video-on-demand file ID.
  // 
  // - **mediaid**: the media file ID.
  // 
  // - **jobId**: the task ID.
  // 
  // example:
  // 
  // {         "vodId": "3e34733b40b9a96ccf5c1ff6f69****",         "mediaid": "eb1861d2c9a842340e989dd56****",         "jobId": "7d2fbc380b0e08e55fe98733764****"     }
  JobInfo *string `json:"JobInfo,omitempty" xml:"JobInfo,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditShowAndReplaceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditShowAndReplaceResponseBody) GoString() string {
  return s.String()
}

func (s *EditShowAndReplaceResponseBody) GetJobInfo() *string  {
  return s.JobInfo
}

func (s *EditShowAndReplaceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditShowAndReplaceResponseBody) SetJobInfo(v string) *EditShowAndReplaceResponseBody {
  s.JobInfo = &v
  return s
}

func (s *EditShowAndReplaceResponseBody) SetRequestId(v string) *EditShowAndReplaceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditShowAndReplaceResponseBody) Validate() error {
  return dara.Validate(s)
}

