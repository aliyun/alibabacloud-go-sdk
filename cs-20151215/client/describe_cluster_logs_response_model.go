// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterLogsResponse
	GetStatusCode() *int32
	SetBody(v []*DescribeClusterLogsResponseBody) *DescribeClusterLogsResponse
	GetBody() []*DescribeClusterLogsResponseBody
}

type DescribeClusterLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DescribeClusterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribeClusterLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterLogsResponse) GetBody() []*DescribeClusterLogsResponseBody {
	return s.Body
}

func (s *DescribeClusterLogsResponse) SetHeaders(v map[string]*string) *DescribeClusterLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterLogsResponse) SetStatusCode(v int32) *DescribeClusterLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterLogsResponse) SetBody(v []*DescribeClusterLogsResponseBody) *DescribeClusterLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterLogsResponse) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterLogsResponseBody struct {
	// The ID of the log entry.
	//
	// example:
	//
	// 590749245
	ID *int64 `json:"ID,omitempty" xml:"ID,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c23421cfa74454bc8b37163fd19af***
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The log content.
	//
	// example:
	//
	// start to update cluster status to CREATE_COMPLETE
	ClusterLog *string `json:"cluster_log,omitempty" xml:"cluster_log,omitempty"`
	// The time when the log entry was generated.
	//
	// example:
	//
	// 2020-09-11T10:11:51+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The time when the log entry was updated.
	//
	// example:
	//
	// 2020-09-11T10:11:51+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeClusterLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterLogsResponseBody) GetID() *int64 {
	return s.ID
}

func (s *DescribeClusterLogsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterLogsResponseBody) GetClusterLog() *string {
	return s.ClusterLog
}

func (s *DescribeClusterLogsResponseBody) GetCreated() *string {
	return s.Created
}

func (s *DescribeClusterLogsResponseBody) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeClusterLogsResponseBody) SetID(v int64) *DescribeClusterLogsResponseBody {
	s.ID = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) SetClusterId(v string) *DescribeClusterLogsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) SetClusterLog(v string) *DescribeClusterLogsResponseBody {
	s.ClusterLog = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) SetCreated(v string) *DescribeClusterLogsResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) SetUpdated(v string) *DescribeClusterLogsResponseBody {
	s.Updated = &v
	return s
}

func (s *DescribeClusterLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
