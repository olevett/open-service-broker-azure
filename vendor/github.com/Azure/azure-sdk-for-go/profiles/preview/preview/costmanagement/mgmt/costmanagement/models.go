// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package costmanagement

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/costmanagement/mgmt/2018-08-01-preview/costmanagement"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AlertCategory = original.AlertCategory

const (
	Billing AlertCategory = original.Billing
	Cost    AlertCategory = original.Cost
	Usage   AlertCategory = original.Usage
)

type AlertCriteria = original.AlertCriteria

const (
	CostThresholdExceeded  AlertCriteria = original.CostThresholdExceeded
	CreditThresholdReached AlertCriteria = original.CreditThresholdReached
	UsageThresholdExceeded AlertCriteria = original.UsageThresholdExceeded
)

type AlertSource = original.AlertSource

const (
	Preset AlertSource = original.Preset
	User   AlertSource = original.User
)

type AlertStatus = original.AlertStatus

const (
	Active     AlertStatus = original.Active
	Dismissed  AlertStatus = original.Dismissed
	Overridden AlertStatus = original.Overridden
	Resolved   AlertStatus = original.Resolved
)

type AlertType = original.AlertType

const (
	Budget  AlertType = original.Budget
	Credit  AlertType = original.Credit
	Invoice AlertType = original.Invoice
)

type ConnectorStatus = original.ConnectorStatus

const (
	ConnectorStatusActive    ConnectorStatus = original.ConnectorStatusActive
	ConnectorStatusError     ConnectorStatus = original.ConnectorStatusError
	ConnectorStatusSuspended ConnectorStatus = original.ConnectorStatusSuspended
)

type ExecutionStatus = original.ExecutionStatus

const (
	Completed  ExecutionStatus = original.Completed
	Failed     ExecutionStatus = original.Failed
	InProgress ExecutionStatus = original.InProgress
	Queud      ExecutionStatus = original.Queud
	Timeout    ExecutionStatus = original.Timeout
)

type ExecutionType = original.ExecutionType

const (
	OnDemand  ExecutionType = original.OnDemand
	Scheduled ExecutionType = original.Scheduled
)

type FormatType = original.FormatType

const (
	Csv FormatType = original.Csv
)

type GranularityType = original.GranularityType

const (
	Daily GranularityType = original.Daily
)

type RecurrenceType = original.RecurrenceType

const (
	RecurrenceTypeAnnually RecurrenceType = original.RecurrenceTypeAnnually
	RecurrenceTypeDaily    RecurrenceType = original.RecurrenceTypeDaily
	RecurrenceTypeMonthly  RecurrenceType = original.RecurrenceTypeMonthly
	RecurrenceTypeWeekly   RecurrenceType = original.RecurrenceTypeWeekly
)

type ReportColumnType = original.ReportColumnType

const (
	ReportColumnTypeDimension ReportColumnType = original.ReportColumnTypeDimension
	ReportColumnTypeTag       ReportColumnType = original.ReportColumnTypeTag
)

type StatusType = original.StatusType

const (
	StatusTypeActive   StatusType = original.StatusTypeActive
	StatusTypeInactive StatusType = original.StatusTypeInactive
)

type TimeframeType = original.TimeframeType

const (
	Custom      TimeframeType = original.Custom
	MonthToDate TimeframeType = original.MonthToDate
	WeekToDate  TimeframeType = original.WeekToDate
)

type Alert = original.Alert
type AlertDefinition = original.AlertDefinition
type AlertListResult = original.AlertListResult
type AlertListResultIterator = original.AlertListResultIterator
type AlertListResultPage = original.AlertListResultPage
type AlertProperties = original.AlertProperties
type AlertsClient = original.AlertsClient
type BaseClient = original.BaseClient
type BillingAccountDimensionsClient = original.BillingAccountDimensionsClient
type CommonReportProperties = original.CommonReportProperties
type ConnectorClient = original.ConnectorClient
type ConnectorCollectionErrorInfo = original.ConnectorCollectionErrorInfo
type ConnectorCollectionInfo = original.ConnectorCollectionInfo
type ConnectorDefinition = original.ConnectorDefinition
type ConnectorDefinitionListResult = original.ConnectorDefinitionListResult
type ConnectorProperties = original.ConnectorProperties
type Dimension = original.Dimension
type DimensionProperties = original.DimensionProperties
type DimensionsListResult = original.DimensionsListResult
type ErrorBase = original.ErrorBase
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Query = original.Query
type QueryColumn = original.QueryColumn
type QueryProperties = original.QueryProperties
type QueryResult = original.QueryResult
type Report = original.Report
type ReportAggregation = original.ReportAggregation
type ReportComparisonExpression = original.ReportComparisonExpression
type ReportDataset = original.ReportDataset
type ReportDatasetConfiguration = original.ReportDatasetConfiguration
type ReportDefinition = original.ReportDefinition
type ReportDeliveryDestination = original.ReportDeliveryDestination
type ReportDeliveryInfo = original.ReportDeliveryInfo
type ReportExecution = original.ReportExecution
type ReportExecutionListResult = original.ReportExecutionListResult
type ReportExecutionProperties = original.ReportExecutionProperties
type ReportFilter = original.ReportFilter
type ReportGrouping = original.ReportGrouping
type ReportListResult = original.ReportListResult
type ReportProperties = original.ReportProperties
type ReportRecurrencePeriod = original.ReportRecurrencePeriod
type ReportSchedule = original.ReportSchedule
type ReportTimePeriod = original.ReportTimePeriod
type ReportsClient = original.ReportsClient
type Resource = original.Resource
type ResourceGroupDimensionsClient = original.ResourceGroupDimensionsClient
type SubscriptionDimensionsClient = original.SubscriptionDimensionsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAlertListResultIterator(page AlertListResultPage) AlertListResultIterator {
	return original.NewAlertListResultIterator(page)
}
func NewAlertListResultPage(getNextPage func(context.Context, AlertListResult) (AlertListResult, error)) AlertListResultPage {
	return original.NewAlertListResultPage(getNextPage)
}
func NewAlertsClient(subscriptionID string) AlertsClient {
	return original.NewAlertsClient(subscriptionID)
}
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBillingAccountDimensionsClient(subscriptionID string) BillingAccountDimensionsClient {
	return original.NewBillingAccountDimensionsClient(subscriptionID)
}
func NewBillingAccountDimensionsClientWithBaseURI(baseURI string, subscriptionID string) BillingAccountDimensionsClient {
	return original.NewBillingAccountDimensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewConnectorClient(subscriptionID string) ConnectorClient {
	return original.NewConnectorClient(subscriptionID)
}
func NewConnectorClientWithBaseURI(baseURI string, subscriptionID string) ConnectorClient {
	return original.NewConnectorClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReportsClient(subscriptionID string) ReportsClient {
	return original.NewReportsClient(subscriptionID)
}
func NewReportsClientWithBaseURI(baseURI string, subscriptionID string) ReportsClient {
	return original.NewReportsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceGroupDimensionsClient(subscriptionID string) ResourceGroupDimensionsClient {
	return original.NewResourceGroupDimensionsClient(subscriptionID)
}
func NewResourceGroupDimensionsClientWithBaseURI(baseURI string, subscriptionID string) ResourceGroupDimensionsClient {
	return original.NewResourceGroupDimensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionDimensionsClient(subscriptionID string) SubscriptionDimensionsClient {
	return original.NewSubscriptionDimensionsClient(subscriptionID)
}
func NewSubscriptionDimensionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionDimensionsClient {
	return original.NewSubscriptionDimensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAlertCategoryValues() []AlertCategory {
	return original.PossibleAlertCategoryValues()
}
func PossibleAlertCriteriaValues() []AlertCriteria {
	return original.PossibleAlertCriteriaValues()
}
func PossibleAlertSourceValues() []AlertSource {
	return original.PossibleAlertSourceValues()
}
func PossibleAlertStatusValues() []AlertStatus {
	return original.PossibleAlertStatusValues()
}
func PossibleAlertTypeValues() []AlertType {
	return original.PossibleAlertTypeValues()
}
func PossibleConnectorStatusValues() []ConnectorStatus {
	return original.PossibleConnectorStatusValues()
}
func PossibleExecutionStatusValues() []ExecutionStatus {
	return original.PossibleExecutionStatusValues()
}
func PossibleExecutionTypeValues() []ExecutionType {
	return original.PossibleExecutionTypeValues()
}
func PossibleFormatTypeValues() []FormatType {
	return original.PossibleFormatTypeValues()
}
func PossibleGranularityTypeValues() []GranularityType {
	return original.PossibleGranularityTypeValues()
}
func PossibleRecurrenceTypeValues() []RecurrenceType {
	return original.PossibleRecurrenceTypeValues()
}
func PossibleReportColumnTypeValues() []ReportColumnType {
	return original.PossibleReportColumnTypeValues()
}
func PossibleStatusTypeValues() []StatusType {
	return original.PossibleStatusTypeValues()
}
func PossibleTimeframeTypeValues() []TimeframeType {
	return original.PossibleTimeframeTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}