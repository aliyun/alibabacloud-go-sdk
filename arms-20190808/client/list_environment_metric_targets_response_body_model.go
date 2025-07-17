// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentMetricTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEnvironmentMetricTargetsResponseBody
	GetCode() *int32
	SetData(v *ListEnvironmentMetricTargetsResponseBodyData) *ListEnvironmentMetricTargetsResponseBody
	GetData() *ListEnvironmentMetricTargetsResponseBodyData
	SetMessage(v string) *ListEnvironmentMetricTargetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvironmentMetricTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEnvironmentMetricTargetsResponseBody
	GetSuccess() *bool
}

type ListEnvironmentMetricTargetsResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	Data *ListEnvironmentMetricTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16AF921B-8187-489F-9913-43C808B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEnvironmentMetricTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentMetricTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentMetricTargetsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEnvironmentMetricTargetsResponseBody) GetData() *ListEnvironmentMetricTargetsResponseBodyData {
	return s.Data
}

func (s *ListEnvironmentMetricTargetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvironmentMetricTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvironmentMetricTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEnvironmentMetricTargetsResponseBody) SetCode(v int32) *ListEnvironmentMetricTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBody) SetData(v *ListEnvironmentMetricTargetsResponseBodyData) *ListEnvironmentMetricTargetsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBody) SetMessage(v string) *ListEnvironmentMetricTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBody) SetRequestId(v string) *ListEnvironmentMetricTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBody) SetSuccess(v bool) *ListEnvironmentMetricTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentMetricTargetsResponseBodyData struct {
	// The active targets.
	ActiveTargets []*ListEnvironmentMetricTargetsResponseBodyDataActiveTargets `json:"ActiveTargets,omitempty" xml:"ActiveTargets,omitempty" type:"Repeated"`
	// The deleted targets.
	DroppedTargets []*ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets `json:"DroppedTargets,omitempty" xml:"DroppedTargets,omitempty" type:"Repeated"`
}

func (s ListEnvironmentMetricTargetsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentMetricTargetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentMetricTargetsResponseBodyData) GetActiveTargets() []*ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	return s.ActiveTargets
}

func (s *ListEnvironmentMetricTargetsResponseBodyData) GetDroppedTargets() []*ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	return s.DroppedTargets
}

func (s *ListEnvironmentMetricTargetsResponseBodyData) SetActiveTargets(v []*ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) *ListEnvironmentMetricTargetsResponseBodyData {
	s.ActiveTargets = v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyData) SetDroppedTargets(v []*ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) *ListEnvironmentMetricTargetsResponseBodyData {
	s.DroppedTargets = v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentMetricTargetsResponseBodyDataActiveTargets struct {
	// The tags used for service discovery.
	DiscoveredLabels map[string]*string `json:"DiscoveredLabels,omitempty" xml:"DiscoveredLabels,omitempty"`
	// The URL of the target.
	//
	// example:
	//
	// http://xxx
	GlobalUrl *string `json:"GlobalUrl,omitempty" xml:"GlobalUrl,omitempty"`
	// The health status.
	//
	// example:
	//
	// up
	Health *string `json:"Health,omitempty" xml:"Health,omitempty"`
	// The tags.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The last error message.
	//
	// example:
	//
	// Get \\"http://172.16.0.86:9104/metrics\\": dial tcp 172.16.0.86:9104: connect: connection refused
	LastError *string `json:"LastError,omitempty" xml:"LastError,omitempty"`
	// The last collection time.
	//
	// example:
	//
	// 2023-10-12T07:15:47.306691514Z
	LastScrape *string `json:"LastScrape,omitempty" xml:"LastScrape,omitempty"`
	// The duration of the last collection.
	//
	// example:
	//
	// 0.00127593
	LastScrapeDuration *float64 `json:"LastScrapeDuration,omitempty" xml:"LastScrapeDuration,omitempty"`
	// The amount of metrics in the last collection.
	//
	// example:
	//
	// 122
	LastScrapeSeries *int64 `json:"LastScrapeSeries,omitempty" xml:"LastScrapeSeries,omitempty"`
	// The name of the collection.
	//
	// example:
	//
	// arms-prom/mysql-exporter-mysql-1694429267986-sm/0"
	ScrapePool *string `json:"ScrapePool,omitempty" xml:"ScrapePool,omitempty"`
	// The URL of the collection.
	//
	// example:
	//
	// http://xxxx
	ScrapeUrl *string `json:"ScrapeUrl,omitempty" xml:"ScrapeUrl,omitempty"`
}

func (s ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GoString() string {
	return s.String()
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetDiscoveredLabels() map[string]*string {
	return s.DiscoveredLabels
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetGlobalUrl() *string {
	return s.GlobalUrl
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetHealth() *string {
	return s.Health
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetLastError() *string {
	return s.LastError
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetLastScrape() *string {
	return s.LastScrape
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetLastScrapeDuration() *float64 {
	return s.LastScrapeDuration
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetLastScrapeSeries() *int64 {
	return s.LastScrapeSeries
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetScrapePool() *string {
	return s.ScrapePool
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) GetScrapeUrl() *string {
	return s.ScrapeUrl
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetDiscoveredLabels(v map[string]*string) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.DiscoveredLabels = v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetGlobalUrl(v string) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.GlobalUrl = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetHealth(v string) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.Health = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetLabels(v map[string]*string) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.Labels = v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetLastError(v string) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.LastError = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetLastScrape(v string) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.LastScrape = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetLastScrapeDuration(v float64) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.LastScrapeDuration = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetLastScrapeSeries(v int64) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.LastScrapeSeries = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetScrapePool(v string) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.ScrapePool = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) SetScrapeUrl(v string) *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets {
	s.ScrapeUrl = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataActiveTargets) Validate() error {
	return dara.Validate(s)
}

type ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets struct {
	// The tags used for service discovery.
	DiscoveredLabels map[string]*string `json:"DiscoveredLabels,omitempty" xml:"DiscoveredLabels,omitempty"`
	// The URL of the target.
	//
	// example:
	//
	// http://xxx
	GlobalUrl *string `json:"GlobalUrl,omitempty" xml:"GlobalUrl,omitempty"`
	// The health status.
	//
	// example:
	//
	// up
	Health *string `json:"Health,omitempty" xml:"Health,omitempty"`
	// The tags.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The last error message.
	//
	// example:
	//
	// Get \\"http://172.16.0.86:9104/metrics\\": dial tcp 172.16.0.86:9104: connect: connection refused
	LastError *string `json:"LastError,omitempty" xml:"LastError,omitempty"`
	// The last collection time.
	//
	// example:
	//
	// 2023-10-12T07:15:47.306691514Z
	LastScrape *string `json:"LastScrape,omitempty" xml:"LastScrape,omitempty"`
	// The duration of the last collection.
	//
	// example:
	//
	// 0.00127593
	LastScrapeDuration *float64 `json:"LastScrapeDuration,omitempty" xml:"LastScrapeDuration,omitempty"`
	// The amount of metrics in the last collection.
	//
	// example:
	//
	// 122
	LastScrapeSeries *int64 `json:"LastScrapeSeries,omitempty" xml:"LastScrapeSeries,omitempty"`
	// The name of the collection.
	//
	// example:
	//
	// arms-prom/mysql-exporter-mysql-1694429267986-sm/0"
	ScrapePool *string `json:"ScrapePool,omitempty" xml:"ScrapePool,omitempty"`
	// The URL of the collection.
	//
	// example:
	//
	// http://xxxx
	ScrapeUrl *string `json:"ScrapeUrl,omitempty" xml:"ScrapeUrl,omitempty"`
}

func (s ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GoString() string {
	return s.String()
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetDiscoveredLabels() map[string]*string {
	return s.DiscoveredLabels
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetGlobalUrl() *string {
	return s.GlobalUrl
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetHealth() *string {
	return s.Health
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetLastError() *string {
	return s.LastError
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetLastScrape() *string {
	return s.LastScrape
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetLastScrapeDuration() *float64 {
	return s.LastScrapeDuration
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetLastScrapeSeries() *int64 {
	return s.LastScrapeSeries
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetScrapePool() *string {
	return s.ScrapePool
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) GetScrapeUrl() *string {
	return s.ScrapeUrl
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetDiscoveredLabels(v map[string]*string) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.DiscoveredLabels = v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetGlobalUrl(v string) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.GlobalUrl = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetHealth(v string) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.Health = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetLabels(v map[string]*string) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.Labels = v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetLastError(v string) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.LastError = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetLastScrape(v string) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.LastScrape = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetLastScrapeDuration(v float64) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.LastScrapeDuration = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetLastScrapeSeries(v int64) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.LastScrapeSeries = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetScrapePool(v string) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.ScrapePool = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) SetScrapeUrl(v string) *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets {
	s.ScrapeUrl = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponseBodyDataDroppedTargets) Validate() error {
	return dara.Validate(s)
}
