import { TableSchema } from '../../services/api/segment'
import { CountriesFormOptions } from '../../lib/countries_timezones'
import { TIMEZONE_OPTIONS } from '../../lib/timezones'
import { Languages } from '../../lib/languages'
import { faUser, faFolderOpen } from '@fortawesome/free-regular-svg-icons'
import { faMousePointer, faBullseye } from '@fortawesome/free-solid-svg-icons'

/**
 * Database table schemas for segmentation engine
 * Based on the actual database structure from internal/database/init.go
 */

export const ContactsTableSchema: TableSchema = {
  name: 'contacts',
  title: 'Contact property',
  description: 'Contact profile and custom fields',
  icon: faUser,
  fields: {
    email: {
      name: 'email',
      title: 'Email',
      description: 'Contact email address',
      type: 'string',
      shown: true
    },
    external_id: {
      name: 'external_id',
      title: 'External ID',
      description: 'External identifier from your system',
      type: 'string',
      shown: true
    },
    first_name: {
      name: 'first_name',
      title: 'First Name',
      description: 'Contact first name',
      type: 'string',
      shown: true
    },
    last_name: {
      name: 'last_name',
      title: 'Last Name',
      description: 'Contact last name',
      type: 'string',
      shown: true
    },
    phone: {
      name: 'phone',
      title: 'Phone',
      description: 'Contact phone number',
      type: 'string',
      shown: true
    },
    country: {
      name: 'country',
      title: 'Country',
      description: 'Contact country',
      type: 'string',
      shown: true,
      options: CountriesFormOptions
    },
    language: {
      name: 'language',
      title: 'Language',
      description: 'Contact language preference',
      type: 'string',
      shown: true,
      options: Languages.map((lang) => ({ value: lang.value, label: lang.name }))
    },
    timezone: {
      name: 'timezone',
      title: 'Timezone',
      description: 'Contact timezone',
      type: 'string',
      shown: true,
      options: TIMEZONE_OPTIONS
    },
    address_line_1: {
      name: 'address_line_1',
      title: 'Address Line 1',
      description: 'Contact address line 1',
      type: 'string',
      shown: true
    },
    address_line_2: {
      name: 'address_line_2',
      title: 'Address Line 2',
      description: 'Contact address line 2',
      type: 'string',
      shown: false
    },
    postcode: {
      name: 'postcode',
      title: 'Postcode',
      description: 'Contact postal code',
      type: 'string',
      shown: true
    },
    state: {
      name: 'state',
      title: 'State',
      description: 'Contact state/province',
      type: 'string',
      shown: true
    },
    job_title: {
      name: 'job_title',
      title: 'Job Title',
      description: 'Contact job title',
      type: 'string',
      shown: true
    },
    // Custom string fields
    custom_string_1: {
      name: 'custom_string_1',
      title: 'Custom String 1',
      description: 'Custom string field 1',
      type: 'string',
      shown: true
    },
    custom_string_2: {
      name: 'custom_string_2',
      title: 'Custom String 2',
      description: 'Custom string field 2',
      type: 'string',
      shown: true
    },
    custom_string_3: {
      name: 'custom_string_3',
      title: 'Custom String 3',
      description: 'Custom string field 3',
      type: 'string',
      shown: true
    },
    custom_string_4: {
      name: 'custom_string_4',
      title: 'Custom String 4',
      description: 'Custom string field 4',
      type: 'string',
      shown: true
    },
    custom_string_5: {
      name: 'custom_string_5',
      title: 'Custom String 5',
      description: 'Custom string field 5',
      type: 'string',
      shown: true
    },
    custom_string_6: {
      name: 'custom_string_6',
      title: 'Custom String 6',
      description: 'Custom string field 6',
      type: 'string',
      shown: true
    },
    custom_string_7: {
      name: 'custom_string_7',
      title: 'Custom String 7',
      description: 'Custom string field 7',
      type: 'string',
      shown: true
    },
    custom_string_8: {
      name: 'custom_string_8',
      title: 'Custom String 8',
      description: 'Custom string field 8',
      type: 'string',
      shown: true
    },
    custom_string_9: {
      name: 'custom_string_9',
      title: 'Custom String 9',
      description: 'Custom string field 9',
      type: 'string',
      shown: true
    },
    custom_string_10: {
      name: 'custom_string_10',
      title: 'Custom String 10',
      description: 'Custom string field 10',
      type: 'string',
      shown: true
    },
    custom_string_11: {
      name: 'custom_string_11',
      title: 'Custom String 11',
      description: 'Custom string field 11',
      type: 'string',
      shown: true
    },
    custom_string_12: {
      name: 'custom_string_12',
      title: 'Custom String 12',
      description: 'Custom string field 12',
      type: 'string',
      shown: true
    },
    custom_string_13: {
      name: 'custom_string_13',
      title: 'Custom String 13',
      description: 'Custom string field 13',
      type: 'string',
      shown: true
    },
    custom_string_14: {
      name: 'custom_string_14',
      title: 'Custom String 14',
      description: 'Custom string field 14',
      type: 'string',
      shown: true
    },
    custom_string_15: {
      name: 'custom_string_15',
      title: 'Custom String 15',
      description: 'Custom string field 15',
      type: 'string',
      shown: true
    },
    // Custom number fields
    custom_number_1: {
      name: 'custom_number_1',
      title: 'Custom Number 1',
      description: 'Custom number field 1',
      type: 'number',
      shown: true
    },
    custom_number_2: {
      name: 'custom_number_2',
      title: 'Custom Number 2',
      description: 'Custom number field 2',
      type: 'number',
      shown: true
    },
    custom_number_3: {
      name: 'custom_number_3',
      title: 'Custom Number 3',
      description: 'Custom number field 3',
      type: 'number',
      shown: true
    },
    custom_number_4: {
      name: 'custom_number_4',
      title: 'Custom Number 4',
      description: 'Custom number field 4',
      type: 'number',
      shown: true
    },
    custom_number_5: {
      name: 'custom_number_5',
      title: 'Custom Number 5',
      description: 'Custom number field 5',
      type: 'number',
      shown: true
    },
    custom_number_6: {
      name: 'custom_number_6',
      title: 'Custom Number 6',
      description: 'Custom number field 6',
      type: 'number',
      shown: true
    },
    custom_number_7: {
      name: 'custom_number_7',
      title: 'Custom Number 7',
      description: 'Custom number field 7',
      type: 'number',
      shown: true
    },
    custom_number_8: {
      name: 'custom_number_8',
      title: 'Custom Number 8',
      description: 'Custom number field 8',
      type: 'number',
      shown: true
    },
    custom_number_9: {
      name: 'custom_number_9',
      title: 'Custom Number 9',
      description: 'Custom number field 9',
      type: 'number',
      shown: true
    },
    custom_number_10: {
      name: 'custom_number_10',
      title: 'Custom Number 10',
      description: 'Custom number field 10',
      type: 'number',
      shown: true
    },
    custom_number_11: {
      name: 'custom_number_11',
      title: 'Custom Number 11',
      description: 'Custom number field 11',
      type: 'number',
      shown: true
    },
    custom_number_12: {
      name: 'custom_number_12',
      title: 'Custom Number 12',
      description: 'Custom number field 12',
      type: 'number',
      shown: true
    },
    custom_number_13: {
      name: 'custom_number_13',
      title: 'Custom Number 13',
      description: 'Custom number field 13',
      type: 'number',
      shown: true
    },
    custom_number_14: {
      name: 'custom_number_14',
      title: 'Custom Number 14',
      description: 'Custom number field 14',
      type: 'number',
      shown: true
    },
    custom_number_15: {
      name: 'custom_number_15',
      title: 'Custom Number 15',
      description: 'Custom number field 15',
      type: 'number',
      shown: true
    },
    // Custom datetime fields
    custom_datetime_1: {
      name: 'custom_datetime_1',
      title: 'Custom Date 1',
      description: 'Custom datetime field 1',
      type: 'time',
      shown: true
    },
    custom_datetime_2: {
      name: 'custom_datetime_2',
      title: 'Custom Date 2',
      description: 'Custom datetime field 2',
      type: 'time',
      shown: true
    },
    custom_datetime_3: {
      name: 'custom_datetime_3',
      title: 'Custom Date 3',
      description: 'Custom datetime field 3',
      type: 'time',
      shown: true
    },
    custom_datetime_4: {
      name: 'custom_datetime_4',
      title: 'Custom Date 4',
      description: 'Custom datetime field 4',
      type: 'time',
      shown: true
    },
    custom_datetime_5: {
      name: 'custom_datetime_5',
      title: 'Custom Date 5',
      description: 'Custom datetime field 5',
      type: 'time',
      shown: true
    },
    custom_datetime_6: {
      name: 'custom_datetime_6',
      title: 'Custom Date 6',
      description: 'Custom datetime field 6',
      type: 'time',
      shown: true
    },
    custom_datetime_7: {
      name: 'custom_datetime_7',
      title: 'Custom Date 7',
      description: 'Custom datetime field 7',
      type: 'time',
      shown: true
    },
    custom_datetime_8: {
      name: 'custom_datetime_8',
      title: 'Custom Date 8',
      description: 'Custom datetime field 8',
      type: 'time',
      shown: true
    },
    custom_datetime_9: {
      name: 'custom_datetime_9',
      title: 'Custom Date 9',
      description: 'Custom datetime field 9',
      type: 'time',
      shown: true
    },
    custom_datetime_10: {
      name: 'custom_datetime_10',
      title: 'Custom Date 10',
      description: 'Custom datetime field 10',
      type: 'time',
      shown: true
    },
    custom_datetime_11: {
      name: 'custom_datetime_11',
      title: 'Custom Date 11',
      description: 'Custom datetime field 11',
      type: 'time',
      shown: true
    },
    custom_datetime_12: {
      name: 'custom_datetime_12',
      title: 'Custom Date 12',
      description: 'Custom datetime field 12',
      type: 'time',
      shown: true
    },
    custom_datetime_13: {
      name: 'custom_datetime_13',
      title: 'Custom Date 13',
      description: 'Custom datetime field 13',
      type: 'time',
      shown: true
    },
    custom_datetime_14: {
      name: 'custom_datetime_14',
      title: 'Custom Date 14',
      description: 'Custom datetime field 14',
      type: 'time',
      shown: true
    },
    custom_datetime_15: {
      name: 'custom_datetime_15',
      title: 'Custom Date 15',
      description: 'Custom datetime field 15',
      type: 'time',
      shown: true
    },
    created_at: {
      name: 'created_at',
      title: 'Created At',
      description: 'Contact creation date',
      type: 'time',
      shown: true
    },
    updated_at: {
      name: 'updated_at',
      title: 'Updated At',
      description: 'Contact last update date',
      type: 'time',
      shown: false
    },
    // Custom JSON fields
    custom_json_1: {
      name: 'custom_json_1',
      title: 'Custom JSON 1',
      description: 'Custom JSON field 1',
      type: 'json',
      shown: true
    },
    custom_json_2: {
      name: 'custom_json_2',
      title: 'Custom JSON 2',
      description: 'Custom JSON field 2',
      type: 'json',
      shown: true
    },
    custom_json_3: {
      name: 'custom_json_3',
      title: 'Custom JSON 3',
      description: 'Custom JSON field 3',
      type: 'json',
      shown: true
    },
    custom_json_4: {
      name: 'custom_json_4',
      title: 'Custom JSON 4',
      description: 'Custom JSON field 4',
      type: 'json',
      shown: true
    },
    custom_json_5: {
      name: 'custom_json_5',
      title: 'Custom JSON 5',
      description: 'Custom JSON field 5',
      type: 'json',
      shown: true
    }
  }
}

export const ContactListsTableSchema: TableSchema = {
  name: 'contact_lists',
  title: 'List subscription',
  description: 'Contact list subscription status',
  icon: faFolderOpen,
  fields: {
    list_id: {
      name: 'list_id',
      title: 'List ID',
      description: 'List identifier',
      type: 'string',
      shown: true
    },
    status: {
      name: 'status',
      title: 'Status',
      description: 'Subscription status',
      type: 'string',
      shown: true,
      options: [
        { value: 'active', label: 'Active' },
        { value: 'unsubscribed', label: 'Unsubscribed' },
        { value: 'pending', label: 'Pending' },
        { value: 'bounced', label: 'Bounced' },
        { value: 'complained', label: 'Complained' }
      ]
    },
    created_at: {
      name: 'created_at',
      title: 'Subscribed At',
      description: 'Date when contact was added to list',
      type: 'time',
      shown: true
    },
    updated_at: {
      name: 'updated_at',
      title: 'Updated At',
      description: 'Last status update date',
      type: 'time',
      shown: false
    },
    deleted_at: {
      name: 'deleted_at',
      title: 'Deleted At',
      description: 'Date when contact was removed from list',
      type: 'time',
      shown: false
    }
  }
}

export const ContactTimelineTableSchema: TableSchema = {
  name: 'contact_timeline',
  title: 'Activity',
  description: 'Contact activity and change history',
  icon: faMousePointer,
  fields: {
    operation: {
      name: 'operation',
      title: 'Operation',
      description: 'Type of operation performed',
      type: 'string',
      shown: true,
      options: [
        { value: 'insert', label: 'Insert' },
        { value: 'update', label: 'Update' }
      ]
    },
    entity_type: {
      name: 'entity_type',
      title: 'Entity Type',
      description: 'Type of entity that changed',
      type: 'string',
      shown: true,
      options: [
        { value: 'contact', label: 'Contact' },
        { value: 'contact_list', label: 'Contact List' },
        { value: 'message_history', label: 'Message History' },
        { value: 'inbound_webhook_event', label: 'Inbound Webhook Event' }
      ]
    },
    entity_id: {
      name: 'entity_id',
      title: 'Entity ID',
      description: 'ID of the related entity',
      type: 'string',
      shown: true
    },
    created_at: {
      name: 'created_at',
      title: 'Event Date',
      description: 'When the event occurred',
      type: 'time',
      shown: true
    }
  }
}

export const CustomEventsGoalsTableSchema: TableSchema = {
  name: 'custom_events_goals',
  title: 'Custom Events Goal',
  description: 'Aggregated custom events data (LTV, transaction counts, etc.)',
  icon: faBullseye,
  fields: {
    goal_type: {
      name: 'goal_type',
      title: 'Goal Type',
      description: 'Type of goal (purchase, subscription, lead, etc.)',
      type: 'string',
      shown: true,
      options: [
        { value: '*', label: 'All types' },
        { value: 'purchase', label: 'Purchase' },
        { value: 'subscription', label: 'Subscription' },
        { value: 'lead', label: 'Lead' },
        { value: 'signup', label: 'Signup' },
        { value: 'booking', label: 'Booking' },
        { value: 'trial', label: 'Trial' },
        { value: 'other', label: 'Other' }
      ]
    },
    goal_name: {
      name: 'goal_name',
      title: 'Goal Name',
      description: 'Optional specific goal name to filter by',
      type: 'string',
      shown: true
    },
    aggregate_operator: {
      name: 'aggregate_operator',
      title: 'Aggregate',
      description: 'How to aggregate the goal values',
      type: 'string',
      shown: true,
      options: [
        { value: 'sum', label: 'Sum' },
        { value: 'count', label: 'Count' },
        { value: 'avg', label: 'Average' },
        { value: 'min', label: 'Minimum' },
        { value: 'max', label: 'Maximum' }
      ]
    },
    operator: {
      name: 'operator',
      title: 'Comparison',
      description: 'Comparison operator',
      type: 'string',
      shown: true,
      options: [
        { value: 'gte', label: 'Greater than or equal' },
        { value: 'lte', label: 'Less than or equal' },
        { value: 'eq', label: 'Equal to' },
        { value: 'between', label: 'Between' }
      ]
    },
    value: {
      name: 'value',
      title: 'Value',
      description: 'Comparison value',
      type: 'number',
      shown: true
    },
    value_2: {
      name: 'value_2',
      title: 'Value 2',
      description: 'Second value for between operator',
      type: 'number',
      shown: true
    }
  }
}

// Export all schemas as a map
export const TableSchemas: { [key: string]: TableSchema } = {
  contacts: ContactsTableSchema,
  contact_lists: ContactListsTableSchema,
  contact_timeline: ContactTimelineTableSchema,
  custom_events_goals: CustomEventsGoalsTableSchema
}
